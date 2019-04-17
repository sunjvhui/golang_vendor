package storage

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/qiniu/api.v7/auth"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/client"
)

var (
	testAK                  = os.Getenv("accessKey")
	testSK                  = os.Getenv("secretKey")
	testBucket              = os.Getenv("QINIU_TEST_BUCKET")
	testBucketPrivate       = os.Getenv("QINIU_TEST_BUCKET_PRIVATE")
	testBucketPrivateDomain = os.Getenv("QINIU_TEST_DOMAIN_PRIVATE")
	testPipeline            = os.Getenv("QINIU_TEST_PIPELINE")

	testKey      = "qiniu.png"
	testFetchUrl = "http://devtools.qiniu.com/qiniu.png"
	testSiteUrl  = "http://devtools.qiniu.com"
)

// 现在qbox.Mac是auth.Credentials的别名， 这个地方使用原来的qbox.Mac
// 测试兼容性是否正确
var (
	mac              *qbox.Mac
	bucketManager    *BucketManager
	operationManager *OperationManager
	formUploader     *FormUploader
	resumeUploader   *ResumeUploader
	base64Uploader   *Base64Uploader
	clt              client.Client
)

func init() {
	clt = client.Client{
		Client: &http.Client{
			Timeout: time.Minute * 10,
		},
	}
	mac = auth.New(testAK, testSK)
	cfg := Config{}
	cfg.Zone = &Zone_z0
	cfg.UseCdnDomains = true
	bucketManager = NewBucketManagerEx(mac, &cfg, &clt)
	operationManager = NewOperationManagerEx(mac, &cfg, &clt)
	formUploader = NewFormUploaderEx(&cfg, &clt)
	resumeUploader = NewResumeUploaderEx(&cfg, &clt)
	base64Uploader = NewBase64UploaderEx(&cfg, &clt)
	rand.Seed(time.Now().Unix())
}

//Test get zone
func TestGetZone(t *testing.T) {
	zone, err := GetZone(testAK, testBucket)
	if err != nil {
		t.Fatalf("GetZone() error, %s", err)
	}
	t.Log(zone.String())
}

// TestCreate 测试创建空间的功能
func TestCreate(t *testing.T) {
	err := bucketManager.CreateBucket("gosdk-test111111111", RIDHuadong)
	if err != nil {
		if err.Error() != "bucket exists" {
			t.Fatalf("CreateBucket() error: %v\n", err)
		}
	}
}

// TestUpdateObjectStatus 测试更新文件状态的功能
func TestUpdateObjectStatus(t *testing.T) {
	keysToStat := []string{"qiniu.png"}

	for _, eachKey := range keysToStat {
		err := bucketManager.UpdateObjectStatus(testBucket, eachKey, false)
		if err != nil {
			if !strings.Contains(err.Error(), "already disabled") {
				t.Fatalf("UpdateObjectStatus error: %v\n", err)
			}
		}
		err = bucketManager.UpdateObjectStatus(testBucket, eachKey, true)
		if err != nil {
			if !strings.Contains(err.Error(), "already enabled") {
				t.Fatalf("UpdateObjectStatus error: %v\n", err)
			}
		}
	}
}

//Test get bucket list
func TestBuckets(t *testing.T) {
	shared := true
	buckets, err := bucketManager.Buckets(shared)
	if err != nil {
		t.Fatalf("Buckets() error, %s", err)
	}

	for _, bucket := range buckets {
		t.Log(bucket)
	}
}

//Test get file info
func TestStat(t *testing.T) {
	keysToStat := []string{"qiniu.png"}

	for _, eachKey := range keysToStat {
		info, err := bucketManager.Stat(testBucket, eachKey)
		if err != nil {
			t.Logf("Stat() error, %s", err)
			t.Fail()
		} else {
			t.Logf("FileInfo:\n %s", info.String())
		}
	}
}

func TestCopyMoveDelete(t *testing.T) {
	keysCopyTarget := []string{"qiniu_1.png", "qiniu_2.png", "qiniu_3.png"}
	keysToDelete := make([]string, 0, len(keysCopyTarget))
	for _, eachKey := range keysCopyTarget {
		err := bucketManager.Copy(testBucket, testKey, testBucket, eachKey, true)
		if err != nil {
			t.Logf("Copy() error, %s", err)
			t.Fail()
		}
	}

	for _, eachKey := range keysCopyTarget {
		keyToDelete := eachKey + "_move"
		err := bucketManager.Move(testBucket, eachKey, testBucket, keyToDelete, true)
		if err != nil {
			t.Logf("Move() error, %s", err)
			t.Fail()
		} else {
			keysToDelete = append(keysToDelete, keyToDelete)
		}
	}

	for _, eachKey := range keysToDelete {
		err := bucketManager.Delete(testBucket, eachKey)
		if err != nil {
			t.Logf("Delete() error, %s", err)
			t.Fail()
		}
	}
}

func TestFetch(t *testing.T) {
	ret, err := bucketManager.Fetch(testFetchUrl, testBucket, "qiniu-fetch.png")
	if err != nil {
		t.Logf("Fetch() error, %s", err)
		t.Fail()
	} else {
		t.Logf("FetchRet:\n %s", ret.String())
	}
}

func TestAsyncFetch(t *testing.T) {

	param := AsyncFetchParam{Url: testFetchUrl, Bucket: testBucket}
	ret, err := bucketManager.AsyncFetch(param)
	if err != nil {
		t.Logf("Fetch() error, %s", err)
		t.Fail()
	} else {
		t.Logf("FetchRet:\n %#v", ret)
	}
}

func TestFetchWithoutKey(t *testing.T) {
	ret, err := bucketManager.FetchWithoutKey(testFetchUrl, testBucket)
	if err != nil {
		t.Logf("FetchWithoutKey() error, %s", err)
		t.Fail()
	} else {
		t.Logf("FetchRet:\n %s", ret.String())
	}
}

func TestDeleteAfterDays(t *testing.T) {
	deleteKey := testKey + "_deleteAfterDays"
	days := 7
	bucketManager.Copy(testBucket, testKey, testBucket, deleteKey, true)
	err := bucketManager.DeleteAfterDays(testBucket, deleteKey, days)
	if err != nil {
		t.Logf("DeleteAfterDays() error, %s", err)
		t.Fail()
	}
}

func TestChangeMime(t *testing.T) {
	toChangeKey := testKey + "_changeMime"
	bucketManager.Copy(testBucket, testKey, testBucket, toChangeKey, true)
	newMime := "text/plain"
	err := bucketManager.ChangeMime(testBucket, toChangeKey, newMime)
	if err != nil {
		t.Fatalf("ChangeMime() error, %s", err)
	}

	info, err := bucketManager.Stat(testBucket, toChangeKey)
	if err != nil || info.MimeType != newMime {
		t.Fatalf("ChangeMime() failed, %s", err)
	}
	bucketManager.Delete(testBucket, toChangeKey)
}

func TestChangeType(t *testing.T) {
	toChangeKey := fmt.Sprintf("%s_changeType_%d", testKey, rand.Int())
	bucketManager.Copy(testBucket, testKey, testBucket, toChangeKey, true)
	fileType := 1
	err := bucketManager.ChangeType(testBucket, toChangeKey, fileType)
	if err != nil {
		t.Fatalf("ChangeType() error, %s", err)
	}

	info, err := bucketManager.Stat(testBucket, toChangeKey)
	if err != nil || info.Type != fileType {
		t.Fatalf("ChangeMime() failed, %s", err)
	}
	bucketManager.Delete(testBucket, toChangeKey)
}

/*
// SetImage成功以后， 后台生效需要一段时间；导致集成测试经常失败。
// 如果要修改这一部分代码可以重新开启这个测试
func TestPrefetchAndImage(t *testing.T) {
	err := bucketManager.SetImage(testSiteUrl, testBucket)
	if err != nil {
		t.Fatalf("SetImage() error, %s", err)
	}

	t.Log("set image success for bucket", testBucket)
	//wait for image set to take effect
	time.Sleep(time.Second * 10)

	err = bucketManager.Prefetch(testBucket, testKey)
	if err != nil {
		t.Fatalf("Prefetch() error, %s", err)
	}

	err = bucketManager.UnsetImage(testBucket)
	if err != nil {
		t.Fatalf("UnsetImage() error, %s", err)
	}

	t.Log("unset image success for bucket", testBucket)
}
*/

func TestListFiles(t *testing.T) {
	limit := 100
	prefix := "listfiles/"
	for i := 0; i < limit; i++ {
		newKey := fmt.Sprintf("%s%s/%d", prefix, testKey, i)
		bucketManager.Copy(testBucket, testKey, testBucket, newKey, true)
	}
	entries, _, _, hasNext, err := bucketManager.ListFiles(testBucket, prefix, "", "", limit)
	if err != nil {
		t.Fatalf("ListFiles() error, %s", err)
	}

	if hasNext {
		t.Fatalf("ListFiles() failed, unexpected hasNext")
	}

	if len(entries) != limit {
		t.Fatalf("ListFiles() failed, unexpected items count, expected: %d, actual: %d", limit, len(entries))
	}

	for _, entry := range entries {
		t.Logf("ListItem:\n%s", entry.String())
	}
}

/*

CDN 节点经常超时
func TestMakePrivateUrl(t *testing.T) {
	deadline := time.Now().Add(time.Second * 3600).Unix()
	privateURL := MakePrivateURL(mac, "http://"+testBucketPrivateDomain, testKey, deadline)
	t.Logf("PrivateUrl: %s", privateURL)
	resp, respErr := clt.Get(privateURL)
	if respErr != nil {
		t.Fatalf("MakePrivateUrl() error, %s", respErr)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("MakePrivateUrl() error, %s", resp.Status)
	}
}
*/

func TestBatch(t *testing.T) {
	copyCnt := 100
	copyOps := make([]string, 0, copyCnt)
	testKeys := make([]string, 0, copyCnt)
	for i := 0; i < copyCnt; i++ {
		cpKey := fmt.Sprintf("%s_batchcopy_%d", testKey, i)
		testKeys = append(testKeys, cpKey)
		copyOps = append(copyOps, URICopy(testBucket, testKey, testBucket, cpKey, true))
	}

	_, bErr := bucketManager.Batch(copyOps)
	if bErr != nil {
		t.Fatalf("BatchCopy error, %s", bErr)
	}

	statOps := make([]string, 0, copyCnt)
	for _, k := range testKeys {
		statOps = append(statOps, URIStat(testBucket, k))
	}
	batchOpRets, bErr := bucketManager.Batch(statOps)
	_, bErr = bucketManager.Batch(copyOps)
	if bErr != nil {
		t.Fatalf("BatchStat error, %s", bErr)
	}

	t.Logf("BatchStat: %v", batchOpRets)
}
