package strtree

import (
	"io"
   "fmt"
   "sort"
   "bytes"
   "errors"
   "strings"
   "reflect"
   "crypto/md5"
   "encoding/gob"
)

type Node struct {
   ch rune
   next []Node
   data interface{}
}

type Value struct {
   Key string
   Data interface{}
}

type Marshaler interface{
	Marshal(io.Writer) error
	Unmarshal(io.Reader) error
	Get() interface{}
}

type NativeMarshaler interface{
	Marshaler
	Set(interface{})
}

var ErrNotFound = errors.New("Not found")
var ErrNeedPointer = errors.New("Gob encoder must be a pointer")
var ErrUnknown = errors.New("Unknown type on data")

var types map[string]Marshaler
var nat2marsh map[string]string


type str struct{
	data string
}
func (t *str) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *str) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *str) Get() interface{} {
	return t.data
}
func (t *str) Set(data interface{}) {
	t.data = data.(string)
}



type abyte struct{
	data []byte
}
func (t *abyte) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *abyte) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *abyte) Get() interface{} {
	return t.data
}
func (t *abyte) Set(data interface{}) {
	t.data = data.([]byte)
}



type ui0 struct{
	data uint
}
func (t *ui0) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *ui0) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *ui0) Get() interface{} {
	return t.data
}
func (t *ui0) Set(data interface{}) {
	t.data = data.(uint)
}



type ui8 struct{
	data uint8
}
func (t *ui8) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *ui8) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *ui8) Get() interface{} {
	return t.data
}
func (t *ui8) Set(data interface{}) {
	t.data = data.(uint8)
}



type ui16 struct{
	data uint16
}
func (t *ui16) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *ui16) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *ui16) Get() interface{} {
	return t.data
}
func (t *ui16) Set(data interface{}) {
	t.data = data.(uint16)
}



type ui32 struct{
	data uint32
}
func (t *ui32) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *ui32) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *ui32) Get() interface{} {
	return t.data
}
func (t *ui32) Set(data interface{}) {
	t.data = data.(uint32)
}



type ui64 struct{
	data uint64
}
func (t *ui64) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *ui64) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *ui64) Get() interface{} {
	return t.data
}
func (t *ui64) Set(data interface{}) {
	t.data = data.(uint64)
}



type i0 struct{
	data int
}
func (t *i0) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *i0) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *i0) Get() interface{} {
	return t.data
}
func (t *i0) Set(data interface{}) {
	t.data = data.(int)
}



type i8 struct{
	data int8
}
func (t *i8) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *i8) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *i8) Get() interface{} {
	return t.data
}
func (t *i8) Set(data interface{}) {
	t.data = data.(int8)
}



type i16 struct{
	data int16
}
func (t *i16) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *i16) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *i16) Get() interface{} {
	return t.data
}
func (t *i16) Set(data interface{}) {
	t.data = data.(int16)
}



type i32 struct{
	data int32
}
func (t *i32) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *i32) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *i32) Get() interface{} {
	return t.data
}
func (t *i32) Set(data interface{}) {
	t.data = data.(int32)
}



type i64 struct{
	data int64
}
func (t *i64) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *i64) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *i64) Get() interface{} {
	return t.data
}
func (t *i64) Set(data interface{}) {
	t.data = data.(int64)
}



type f32 struct{
	data float32
}
func (t *f32) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *f32) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *f32) Get() interface{} {
	return t.data
}
func (t *f32) Set(data interface{}) {
	t.data = data.(float32)
}



type f64 struct{
	data float64
}
func (t *f64) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *f64) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *f64) Get() interface{} {
	return t.data
}
func (t *f64) Set(data interface{}) {
	t.data = data.(float64)
}



type b struct{
	data bool
}
func (t *b) Unmarshal(rd io.Reader) error {
	return gob.NewDecoder(rd).Decode(&t.data)
}
func (t *b) Marshal(w io.Writer) error {
	return gob.NewEncoder(w).Encode(t.data)
}
func (t *b) Get() interface{} {
	return t.data
}
func (t *b) Set(data interface{}) {
	t.data = data.(bool)
}




func init() {
	var err error
	nat2marsh = map[string]string{}

	err = Register(&str{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId("")] = genId(&str{})

	err = Register(&abyte{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId([]byte{})] = genId(&abyte{})

	err = Register(&ui0{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(uint(0))] = genId(&ui0{})

	err = Register(&ui8{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(uint8(0))] = genId(&ui8{})

	err = Register(&ui16{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(uint16(0))] = genId(&ui16{})

	err = Register(&ui32{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(uint32(0))] = genId(&ui32{})

	err = Register(&ui64{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(uint64(0))] = genId(&ui64{})

	err = Register(&i0{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(int(0))] = genId(&i0{})

	err = Register(&i8{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(int8(0))] = genId(&i8{})

	err = Register(&i16{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(int16(0))] = genId(&i16{})

	err = Register(&i32{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(int32(0))] = genId(&i32{})

	err = Register(&i64{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(int64(0))] = genId(&i64{})

	err = Register(&f32{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(float32(0))] = genId(&f32{})

	err = Register(&f64{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(float64(0))] = genId(&f64{})

	err = Register(&b{})
	if err != nil {
		panic(err)
	}
	nat2marsh[genId(true)] = genId(&b{})

}


func genId(typ interface{}) string {
	var hsh [md5.Size]byte
	var i int
	var id []byte
	var b byte

	hsh = md5.Sum([]byte(reflect.TypeOf(typ).String()))
	id = make([]byte,8)

	for i, b = range hsh {
		id[i&0x7] ^= b
	} 
	
	return string(id)
}

func Register(m Marshaler) error {
	var v reflect.Value

	v = reflect.ValueOf(m)
	if v.Kind() != reflect.Pointer {
		return ErrNeedPointer
	}

	if len(types) == 0 {
		types = map[string]Marshaler{}
	}

	types[genId(m)] = m
	return nil
}

func (n *Node) GobDecode(input []byte) error {
	var buf *bytes.Buffer
	var buf2 string
	var err error
	var dec *gob.Decoder
	var m Marshaler
	var nm NativeMarshaler
	var ok bool

	buf = bytes.NewBuffer(input)
	n.ch, _, err = buf.ReadRune()
	if err != nil {
		return err
	}

	dec = gob.NewDecoder(buf)
	err = dec.Decode(&n.next)
	if err != nil && fmt.Sprintf("%s", err) != "EOF" {
		return err
	}

	err = dec.Decode(&buf2)
	if err != nil && fmt.Sprintf("%s", err) != "EOF" {
		return err
	}

//	fmt.Printf("### %#v\n", buf2)

	if buf2 != "X" {
		m, ok = types[string(buf2)]
		if !ok {
			return err
		}

		err = m.Unmarshal(buf)
		if err != nil && fmt.Sprintf("%s", err) != "EOF" {
			return ErrUnknown
		}

		nm, ok = m.(NativeMarshaler)
		if ok {
			n.data = nm.Get()
		} else {
			n.data = m
		}
	}

//	fmt.Printf(">>>>>>>> i=%#v\n", n.data)

	return nil
}

func (n *Node) GobEncode() ([]byte, error) {
	var buf *bytes.Buffer
	var err error
	var enc *gob.Encoder
	var m Marshaler
	var nm NativeMarshaler
	var id string
	var ok bool

	buf = bytes.NewBuffer([]byte{})
	buf.WriteRune(n.ch)

	enc = gob.NewEncoder(buf)
	err = enc.Encode(n.next)
	if err != nil {
		return buf.Bytes(), err
	}

	if n.data != nil {
		id, ok = nat2marsh[genId(n.data)]
		if ok {
			err = enc.Encode(id)
			if err != nil {
				return buf.Bytes(), err
			}

			m = types[id]
			nm, ok = m.(NativeMarshaler)
			if !ok {
				return buf.Bytes(), ErrUnknown
			}
			nm.Set(n.data)
			err = nm.Marshal(buf)
			if err != nil {
				return buf.Bytes(), err
			}
		} else {
			m, ok = n.data.(Marshaler)
			if ok {
				err = enc.Encode(genId(n.data))
				if err != nil {
					return buf.Bytes(), err
				}

				err = m.Marshal(buf)
				if err != nil {
					return buf.Bytes(), err
				}
			}
		}
	} else {
		err = enc.Encode("X")
		if err != nil {
			return buf.Bytes(), err
		}
	}

	return buf.Bytes(), nil
}

func (n *Node) Fetch(s string) (*Node, int, int, error) {
   var c rune
   var i int
   var curr *Node
   var pos int

   curr = n
   for pos, c = range s {
      i = sort.Search(len(curr.next), func(j int) bool { return curr.next[j].ch >= c })
      if i < len(curr.next) && curr.next[i].ch == c {
         // x is present at data[i]
         curr = &curr.next[i]
      } else {
         // x is not present in data,
         // but i is the index where it would be inserted.
         return curr, i,pos,  ErrNotFound
      }
   }

   return curr, i, pos, nil
}

func (n *Node) Get(s string) (interface{}, error) {
   var err error
   var n2 *Node

   n2, _, _, err = n.Fetch(s)
   if err != nil {
      return nil, err
   }

   if n2.data == nil {
      return nil, ErrNotFound
   }

   return n2.data, nil
}

func (n *Node) Set(s string, val interface{}) error {
   var i, pos int
   var err error
   var n2 *Node
   var c rune

   n2, i, pos, err = n.Fetch(s)

   if err != nil {
      if err != ErrNotFound {
         return err
      }

      for pos, c = range s[pos:] {
         if pos == 0 {
            if len(n2.next) == 0 {
               n2.next = []Node{Node{ch: c, next:[]Node{}}}
            } else {
               if i==len(n2.next) {
                  n2.next = append(n2.next, Node{ch: c, next:[]Node{}})
               } else {
                  n2.next = append(n2.next, Node{next:[]Node{}})
                  copy(n2.next[i+1:],n2.next[i:])
                  n2.next[i].ch = c
               }
            }
            n2 = &n2.next[i]
         } else {
            n2.next = []Node{Node{ch: c, next:[]Node{}}}
            n2 = &n2.next[0]
         }
      }
   }

   n2.data = val

   return nil
}

func (n Node) astring(s string) []string {
   var as []string
   var n2 Node

   if n.data != nil {
      as = []string{fmt.Sprintf("%s%c=%q",s,n.ch,n.data)}
   } else {
      as = []string{}
   }

   if n.ch == 0 {
      return as
   }

   for _, n2 = range n.next {
      as = append(as,n2.astring(fmt.Sprintf("%s%c",s,n.ch))...)
   }

   return as
}

func (kv Value) String() string {
	return fmt.Sprintf("%s: %#v", kv.Key, kv.Value)
}

func (n Node) String() string {
   var s, sep string
   var n2 Node

   if n.ch != 0 {
      s = fmt.Sprintf("%c",n.ch)
   }

   sep = ", " + s
   for _, n2 = range n.next {
      s += strings.Join(n2.astring(s),sep)
   }

   return s
}

func (n Node) list(s string) []Value {
   var n2 Node
   var av []Value

   if n.data != nil {
      av = []Value{Value{Key:fmt.Sprintf("%s%c",s,n.ch),Data: n.data}}
   } else {
      av = []Value{}
   }

   if n.ch != 0 {
      s = fmt.Sprintf("%s%c", s, n.ch)
   }

   for _, n2 = range n.next {
      av = append(av,n2.list(s)...)
   }

   return av
}

func (n Node) List() []Value {
   return n.list("")
}

