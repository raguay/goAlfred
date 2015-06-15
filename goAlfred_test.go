package goAlfred

import (
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func Test_sort_comparison(t *testing.T) {
	Convey("Given two AlfredResults with differing priorities", t, func() {
		a := AlfredResult{Priority: 1, Arg: "A"}
		b := AlfredResult{Priority: 2, Arg: "B"}

		var results ByPriority
		results = ByPriority{a, b}
		Convey("The second is less than the first", func() {
			So(results.Less(0, 1), ShouldBeFalse)
		})

		Convey("The first is greater than the second", func() {
			So(results.Less(1, 0), ShouldBeTrue)
		})
	})
}

func Test_sort_manipulation(t *testing.T) {
	Convey("Given three AlfredResults in a ByPriority slice", t, func() {
		a := AlfredResult{Priority: 1, Arg: "A"}
		b := AlfredResult{Priority: 2, Arg: "B"}
		c := AlfredResult{Priority: 2, Arg: "C"}

		var results ByPriority
		results = ByPriority{a, b, c}

		Convey("Initial order is A, B, C", func() {
			So(results[0].Arg, ShouldEqual, "A")
			So(results[1].Arg, ShouldEqual, "B")
			So(results[2].Arg, ShouldEqual, "C")
		})
		Convey("Swapping the first and the second switches their place", func() {
			results.Swap(0, 1)
			So(results[0].Arg, ShouldEqual, "B")
			So(results[1].Arg, ShouldEqual, "A")
			So(results[2].Arg, ShouldEqual, "C")
		})
		Convey("Swapping the second and third switches their place", func() {
			results.Swap(1, 2)
			So(results[0].Arg, ShouldEqual, "A")
			So(results[1].Arg, ShouldEqual, "C")
			So(results[2].Arg, ShouldEqual, "B")
		})
		Convey("Swapping the first and third switches their place", func() {
			results.Swap(0, 2)
			So(results[0].Arg, ShouldEqual, "C")
			So(results[1].Arg, ShouldEqual, "B")
			So(results[2].Arg, ShouldEqual, "A")
		})
	})
}

func Test_sort_length(t *testing.T) {
	Convey("Given three AlfredResults in a ByPriority  slice", t, func() {
		a := AlfredResult{Priority: 1, Arg: "A"}
		b := AlfredResult{Priority: 2, Arg: "B"}
		c := AlfredResult{Priority: 2, Arg: "C"}

		var results ByPriority
		results = ByPriority{a, b, c}

		Convey("Length is 3", func() {
			So(results.Len(), ShouldEqual, 3)
		})
	})
}

func Test_sort(t *testing.T) {
	Convey("Given a mix of AlfredResults with varying priorities", t, func() {
		a := AlfredResult{Priority: 1, Arg: "A"}
		b := AlfredResult{Priority: 0, Arg: "B"}
		c := AlfredResult{Priority: 5, Arg: "C"}
		d := AlfredResult{Priority: 0, Arg: "D"}
		e := AlfredResult{Priority: 3, Arg: "E"}
		f := AlfredResult{Priority: 2, Arg: "F"}

		var results ByPriority
		results = ByPriority{a, b, c, d, e, f}

		Convey("sort.Sort() sorts by priority", func() {
			sort.Sort(results)
			So(results[0].Arg, ShouldEqual, "C")
			So(results[1].Arg, ShouldEqual, "E")
			So(results[2].Arg, ShouldEqual, "F")
			So(results[3].Arg, ShouldEqual, "A")
			So(results[4].Arg, ShouldEqual, "B")
			So(results[5].Arg, ShouldEqual, "D")
		})
	})
}

func Test_adding_results(t *testing.T) {
	Convey("A result slice with AlfredResults having varying priorities", t, func() {
		results = []AlfredResult{
			AlfredResult{Priority: 1, Arg: "A", Uid: "A"},
			AlfredResult{Priority: 0, Arg: "B", Uid: "B"},
			AlfredResult{Priority: 5, Arg: "C", Uid: "C"},
			AlfredResult{Priority: 0, Arg: "D", Uid: "D"},
			AlfredResult{Priority: 3, Arg: "E", Uid: "E"},
			AlfredResult{Priority: 2, Arg: "F", Uid: "F"},
		}
		maxResults = 6

		Convey("Causes ToXML to sort by priority", func() {
			So(ToXML(), ShouldEqual,
				"<items><item uidid=\"C\" valid=\"\" auto=\"\"><arg>C</arg><title></title><subtitle></subtitle><icon></icon><omit>5</omit></item><item uidid=\"E\" valid=\"\" auto=\"\"><arg>E</arg><title></title><subtitle></subtitle><icon></icon><omit>3</omit></item><item uidid=\"F\" valid=\"\" auto=\"\"><arg>F</arg><title></title><subtitle></subtitle><icon></icon><omit>2</omit></item><item uidid=\"A\" valid=\"\" auto=\"\"><arg>A</arg><title></title><subtitle></subtitle><icon></icon><omit>1</omit></item><item uidid=\"B\" valid=\"\" auto=\"\"><arg>B</arg><title></title><subtitle></subtitle><icon></icon><omit>0</omit></item><item uidid=\"D\" valid=\"\" auto=\"\"><arg>D</arg><title></title><subtitle></subtitle><icon></icon><omit>0</omit></item></items>\n",
			)
		})
	})
}
