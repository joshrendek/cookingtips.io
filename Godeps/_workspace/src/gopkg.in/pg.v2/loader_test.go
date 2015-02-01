package pg_test

import (
	"errors"

	. "launchpad.net/gocheck"

	"gopkg.in/pg.v2"
)

type LoaderTest struct {
	db *pg.DB
}

var _ = Suite(&LoaderTest{})

func (t *LoaderTest) SetUpTest(c *C) {
	t.db = pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "test",
	})
}

func (t *LoaderTest) TearDownTest(c *C) {
	c.Assert(t.db.Close(), IsNil)
}

type numLoader struct {
	Num int
}

func (l *numLoader) New() interface{} {
	return l
}

type numNum2Loader struct {
	*numLoader
	Num2 int
}

func (l *numNum2Loader) New() interface{} {
	return l
}

func (t *LoaderTest) TestQuery(c *C) {
	dst := &numLoader{}
	_, err := t.db.Query(dst, "SELECT 1 AS num")
	c.Assert(err, IsNil)
	c.Assert(dst.Num, Equals, 1)
}

func (t *LoaderTest) TestQueryEmbeddedStruct(c *C) {
	dst := &numNum2Loader{
		numLoader: &numLoader{},
	}
	_, err := t.db.Query(dst, "SELECT 1 AS num, 2 as num2")
	c.Assert(err, IsNil)
	c.Assert(dst.Num, Equals, 1)
	c.Assert(dst.Num2, Equals, 2)
}

func (t *LoaderTest) TestQueryStmt(c *C) {
	stmt, err := t.db.Prepare("SELECT 1 AS num")
	c.Assert(err, IsNil)
	defer stmt.Close()

	dst := &numLoader{}
	_, err = stmt.Query(dst)
	c.Assert(err, IsNil)
	c.Assert(dst.Num, Equals, 1)
}

func (t *LoaderTest) TestQueryInts(c *C) {
	var ids pg.Ints
	_, err := t.db.Query(&ids, "SELECT s.num AS num FROM generate_series(0, 10) AS s(num)")
	c.Assert(err, IsNil)
	c.Assert(ids, DeepEquals, pg.Ints{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func (t *LoaderTest) TestQueryInts2(c *C) {
	var ints pg.Ints
	_, err := t.db.Query(&ints, "SELECT * FROM generate_series(1, 1000000)")
	c.Assert(err, IsNil)
	c.Assert(ints, HasLen, 1000000)
}

func (t *LoaderTest) TestQueryStrings(c *C) {
	var strings pg.Strings
	_, err := t.db.Query(&strings, "SELECT 'hello'")
	c.Assert(err, IsNil)
	c.Assert(strings, DeepEquals, pg.Strings{"hello"})
}

type errLoader string

func (l errLoader) Load(colIdx int, colName string, b []byte) error {
	return errors.New(string(l))
}

func (t *LoaderTest) TestLoaderError(c *C) {
	tx, err := t.db.Begin()
	c.Assert(err, IsNil)
	defer tx.Rollback()

	loader := errLoader("my error")
	_, err = tx.QueryOne(loader, "SELECT 1, 2")
	c.Assert(err.Error(), Equals, "my error")

	// Verify that client is still functional.
	var n1, n2 int
	_, err = tx.QueryOne(pg.LoadInto(&n1, &n2), "SELECT 1, 2")
	c.Assert(err, IsNil)
	c.Assert(n1, Equals, 1)
	c.Assert(n2, Equals, 2)
}

func (t *LoaderTest) BenchmarkQueryRow(c *C) {
	dst := &numLoader{}
	for i := 0; i < c.N; i++ {
		_, err := t.db.QueryOne(dst, "SELECT ?::bigint AS num", 1)
		if err != nil {
			panic(err)
		}
		if dst.Num != 1 {
			panic("dst.Num != 1")
		}
	}
}

func (t *LoaderTest) BenchmarkQueryRowWithoutParams(c *C) {
	dst := &numLoader{}
	for i := 0; i < c.N; i++ {
		_, err := t.db.QueryOne(dst, "SELECT 1::bigint AS num")
		if err != nil {
			panic(err)
		}
		if dst.Num != 1 {
			panic("dst.Num != 1")
		}
	}
}

func (t *LoaderTest) BenchmarkQueryRowStmt(c *C) {
	stmt, err := t.db.Prepare("SELECT $1::bigint AS num")
	c.Assert(err, IsNil)
	defer stmt.Close()

	c.ResetTimer()

	dst := &numLoader{}
	for i := 0; i < c.N; i++ {
		_, err := stmt.QueryOne(dst, 1)
		if err != nil {
			panic(err)
		}
	}
}
