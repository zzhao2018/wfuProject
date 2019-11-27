package midware

import (
	"context"
	"fmt"
	"github.com/MXi4oyu/golang.org/x/time/rate"
	"math"
	"time"
)

//限流
type LimitMid interface {
	Allow() bool
}

/********************计数器限流器******************/
type CounterLimit struct {
	timeUnixNano int64
	reqNum int64
	reqLimit int64
	timeDiff time.Duration
}

func NewCounterLimit(limit int64,diff time.Duration)*CounterLimit{
	return &CounterLimit{
		timeUnixNano: time.Now().UnixNano(),
		reqNum:   0,
		reqLimit:limit,
		timeDiff:diff,
	}
}

func(g *CounterLimit)Allow()bool{
	//计算当前时间与起始时间之差
	timeNowUnix:=time.Now().UnixNano()
	//回复
	if timeNowUnix-g.timeUnixNano>int64(g.timeDiff) {
		g.reqNum=1
		g.timeUnixNano=timeNowUnix
		return true
	}
	//判断是否可以访问
	if g.reqNum>g.reqLimit {
		return false
	}
	g.reqNum++
	return true
}


/********************漏桶限流器*******************/
type BucketLimit struct {
	qbs float64  //速率
	waterCur float64 //漏桶水面
	bucketSize float64 //漏桶大小
	timeCur int64//当前时间
}

func NewBucketLimit(qbs float64,bucketSize float64)*BucketLimit{
	return &BucketLimit{
		qbs:        qbs,
		waterCur:   0,
		bucketSize: bucketSize,
		timeCur:    time.Now().UnixNano(),
	}
}

func(l *BucketLimit)Allow()bool{
	//更新桶
	timeN:=time.Now().UnixNano()
	timeDiff:=float64(timeN-l.timeCur)/1000/1000/1000
	//计算剩余水面
	l.waterCur=math.Max(0,l.waterCur-(l.qbs*timeDiff))
	l.timeCur=timeN
	//判断水面是否溢出
	if l.waterCur>l.bucketSize {
		return false
	}
	l.waterCur++
	return true
}

/********************令牌桶限流器*****************/
type LPLimit struct {
	limiter *rate.Limiter
}

func NewLPLimit(l rate.Limit,size int)*LPLimit{
	return &LPLimit{limiter:rate.NewLimiter(l,size)}
}

func(l *LPLimit)Allow()bool{
	return l.limiter.Allow()
}

/******************生成中间件*****************/
func NewLimitMidware(limit LimitMid)MidWare{
	return func(next MidWareFunc) MidWareFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			allow:=limit.Allow()
			if allow==false {
				err=fmt.Errorf("request out of limit")
				return
			}
			resp,err=next(ctx,req)
			return
		}
	}
}