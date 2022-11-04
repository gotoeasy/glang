package cmn

import (
	"errors"
	"os"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// 存储结构体
type LevelDB struct {
	dbPath  string      // 数据存放目录
	leveldb *leveldb.DB // leveldb
	opened  bool        // 是否打开状态
	mu      sync.Mutex  // 锁
}

// 创建存储对象，参数dbPath为数据库名目录
func NewLevelDB(dbPath string) (*LevelDB, error) {
	db := &LevelDB{
		dbPath: dbPath,
	}
	err := os.MkdirAll(db.dbPath, 0777)
	if err != nil {
		return nil, err
	}
	return db, db.open()
}

// 保存
func (s *LevelDB) Put(key []byte, value []byte) error {
	if !s.opened {
		return errors.New("数据库没有打开")
	}
	return s.leveldb.Put(key, value, nil)
}

// 删除
func (s *LevelDB) Del(key []byte) error {
	if !s.opened {
		return errors.New("数据库没有打开")
	}
	return s.leveldb.Delete(key, nil)
}

// 获取
func (s *LevelDB) Get(key []byte) ([]byte, error) {
	if !s.opened {
		return nil, errors.New("数据库没有打开")
	}
	return s.leveldb.Get(key, nil)
}

// 快照
func (s *LevelDB) GetSnapshot() (*leveldb.Snapshot, error) {
	if !s.opened {
		return nil, errors.New("数据库没有打开")
	}
	return s.leveldb.GetSnapshot()
}

// 打开数据库
func (s *LevelDB) open() error {
	if s.opened {
		return nil
	}

	s.mu.Lock()         // 锁
	defer s.mu.Unlock() // 解锁
	if s.opened {
		return nil
	}

	option := new(opt.Options)                    // leveldb选项
	option.Filter = filter.NewBloomFilter(10)     // 使用布隆过滤器
	db, err := leveldb.OpenFile(s.dbPath, option) // 打开数据库
	if err != nil {
		Error("打开数据库失败：", s.dbPath)
		return err
	}
	s.leveldb = db

	s.opened = true
	Info("打开数据库")
	return nil
}

// 关闭数据库
func (s *LevelDB) Close() {
	if !s.opened {
		return
	}

	s.mu.Lock()         // 锁
	defer s.mu.Unlock() // 解锁
	if !s.opened {
		return
	}

	s.opened = false
	err := s.leveldb.Close()
	if err != nil {
		Error("关闭数据库失败", err)
	} else {
		Info("关闭数据库")
	}

}

// 是否关闭状态
func (s *LevelDB) IsOpen() bool {
	return s.opened
}
