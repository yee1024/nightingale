package istr

import (
	"bytes"
	"sync"
)

const SEPERATOR = "/"

var bufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

//strs目前有三种类别
// endpoint/metric/tags
// endpoint/counter
// metric/tags
// strs 参数必须按照上面的顺序来入参
// func PK(strs ...string) string {
// 	ret := bufferPool.Get().(*bytes.Buffer)
// 	ret.Reset()
// 	defer bufferPool.Put(ret)
// 	count := len(strs)
// 	if count == 0 {
// 		return ""
// 	}

// 	ret.WriteString(strs[0])
// 	for i := 1; i < count-1; i++ {
// 		ret.WriteString(SEPERATOR)
// 		ret.WriteString(strs[i])
// 	}

// 	if strs[count-1] != "" {
// 		ret.WriteString(SEPERATOR)
// 		ret.WriteString(strs[count-1])
// 	}

// 	return ret.String()
// }

// func UUID(endpoint, metric, tags, dstype string, step int) string {
// 	ret := bufferPool.Get().(*bytes.Buffer)
// 	ret.Reset()
// 	defer bufferPool.Put(ret)

// 	if tags == "" {
// 		ret.WriteString(endpoint)
// 		ret.WriteString(SEPERATOR)
// 		ret.WriteString(metric)
// 		ret.WriteString(SEPERATOR)
// 		ret.WriteString(dstype)
// 		ret.WriteString(SEPERATOR)
// 		ret.WriteString(strconv.Itoa(step))

// 		return ret.String()
// 	}
// 	ret.WriteString(endpoint)
// 	ret.WriteString(SEPERATOR)
// 	ret.WriteString(metric)
// 	ret.WriteString(SEPERATOR)
// 	ret.WriteString(tags)
// 	ret.WriteString(SEPERATOR)
// 	ret.WriteString(dstype)
// 	ret.WriteString(SEPERATOR)
// 	ret.WriteString(strconv.Itoa(step))

// 	return ret.String()
// }

// func XXhash(strs ...string) uint64 {
// 	ret := bufferPool.Get().(*bytes.Buffer)
// 	ret.Reset()
// 	defer bufferPool.Put(ret)
// 	count := len(strs)
// 	if count == 0 {
// 		return 0
// 	}

// 	ret.WriteString(strs[0])
// 	for i := 1; i < count-1; i++ {
// 		ret.WriteString(SEPERATOR)
// 		ret.WriteString(strs[i])
// 	}

// 	if strs[count-1] != "" {
// 		ret.WriteString(SEPERATOR)
// 		ret.WriteString(strs[count-1])
// 	}

// 	return xxhash.Sum64(ret.Bytes())
// }

// func ToMD5(endpoint string, metric string, tags string) string {
// 	return str.MD5(PK(endpoint, metric, tags))
// }

// func SortedTags(tags map[string]string) string {
// 	if tags == nil {
// 		return ""
// 	}

// 	size := len(tags)

// 	if size == 0 {
// 		return ""
// 	}

// 	ret := bufferPool.Get().(*bytes.Buffer)
// 	ret.Reset()
// 	defer bufferPool.Put(ret)

// 	if size == 1 {
// 		for k, v := range tags {
// 			ret.WriteString(k)
// 			ret.WriteString("=")
// 			ret.WriteString(v)
// 		}
// 		return ret.String()
// 	}

// 	keys := make([]string, size)
// 	i := 0
// 	for k := range tags {
// 		keys[i] = k
// 		i++
// 	}

// 	sort.Strings(keys)

// 	for j, key := range keys {
// 		ret.WriteString(key)
// 		ret.WriteString("=")
// 		ret.WriteString(tags[key])
// 		if j != size-1 {
// 			ret.WriteString(",")
// 		}
// 	}

// 	return ret.String()
// }

// func UnixTsFormat(ts int64) string {
// 	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
// }

// func IdsString(ids []int) string {
// 	count := len(ids)
// 	arr := make([]string, count)
// 	for i := 0; i < count; i++ {
// 		arr[i] = fmt.Sprintf("%d", ids[i])
// 	}
// 	return strings.Join(arr, ",")
// }
