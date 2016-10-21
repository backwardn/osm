package osm

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
	"time"
)

func TestWayMarshalXML(t *testing.T) {
	w := Way{
		ID: 123,
	}

	data, err := xml.Marshal(w)
	if err != nil {
		t.Fatalf("xml marshal error: %v", err)
	}

	expected := `<way id="123" user="" uid="0" visible="false" version="0" changeset="0" timestamp="0001-01-01T00:00:00Z"></way>`
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf("incorrect marshal, got: %s", string(data))
	}

	// node refs
	w.Nodes = []WayNode{{ID: 123}}
	data, err = xml.Marshal(w)
	if err != nil {
		t.Fatalf("xml marshal error: %v", err)
	}

	if !bytes.Equal(data, []byte(`<way id="123" user="" uid="0" visible="false" version="0" changeset="0" timestamp="0001-01-01T00:00:00Z"><nd ref="123"></nd></way>`)) {
		t.Errorf("not marshalled correctly: %s", string(data))
	}

	// node with lat/lon
	w.Nodes[0] = WayNode{Lat: 1, Lon: 2}
	data, err = xml.Marshal(w)
	if err != nil {
		t.Fatalf("xml marshal error: %v", err)
	}

	if !bytes.Equal(data, []byte(`<way id="123" user="" uid="0" visible="false" version="0" changeset="0" timestamp="0001-01-01T00:00:00Z"><nd ref="0" lat="1" lon="2"></nd></way>`)) {
		t.Errorf("not marshalled correctly: %s", string(data))
	}

	// minor way
	w.Nodes = nil
	w.Minors = []MinorWay{
		MinorWay{
			Timestamp: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
			MinorNodes: []MinorWayNode{
				{Index: 0, Version: 2},
			},
		},
	}
	data, err = xml.Marshal(w)
	if err != nil {
		t.Fatalf("xml marshal error: %v", err)
	}

	if !bytes.Equal(data, []byte(`<way id="123" user="" uid="0" visible="false" version="0" changeset="0" timestamp="0001-01-01T00:00:00Z"><minor-way timestamp="2012-01-01T00:00:00Z"><minor-nd index="0" version="2"></minor-nd></minor-way></way>`)) {
		t.Errorf("not marshalled correctly: %s", string(data))
	}

	// blanket xml test
	data = readFile(t, "testdata/minor-way.osm")

	osm := &OSM{}
	err = xml.Unmarshal(data, &osm)
	if err != nil {
		t.Errorf("unmarshal error: %v", err)
	}
	way := osm.Ways[0]

	var i1 interface{}
	err = xml.Unmarshal(data, &i1)
	if err != nil {
		t.Errorf("unmarshal error: %v", err)
	}

	data, err = xml.Marshal(way)
	if err != nil {
		t.Errorf("marshal error: %v", err)
	}

	var i2 interface{}
	err = xml.Unmarshal(data, &i2)
	if err != nil {
		t.Errorf("unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(i1, i2) {
		t.Errorf("interfaces not equal")
		t.Logf("%+v", i1)
		t.Logf("%+v", i2)
	}
}

func TestWaysMarshal(t *testing.T) {
	ws := Ways{
		{ID: 123},
		{ID: 321},
	}

	data, err := ws.Marshal()
	if err != nil {
		t.Fatalf("ways marshal error: %v", err)
	}

	ws2, err := UnmarshalWays(data)
	if err != nil {
		t.Fatalf("ways unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(ws, ws2) {
		t.Errorf("ways not equal")
		t.Logf("%+v", ws)
		t.Logf("%+v", ws2)
	}

	// empty ways
	ws = Ways{}

	data, err = ws.Marshal()
	if err != nil {
		t.Fatalf("ways marshal error: %v", err)
	}

	if l := len(data); l != 0 {
		t.Errorf("length of way data should be 0, got %v", l)
	}

	ws2, err = UnmarshalWays(data)
	if err != nil {
		t.Fatalf("ways unmarshal error: %v", err)
	}

	if ws2 != nil {
		t.Errorf("should return nil Ways for empty data, got %v", ws2)
	}
}