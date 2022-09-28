package strtree

import (
   "fmt"
   "sort"
   "bytes"
   "errors"
   "strings"
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

var ErrNotFound = errors.New("Not found")

func (n *Node) GobDecode(input []byte) error {
	var buf *bytes.Buffer
	var err error
	var dec *gob.Decoder

	buf = bytes.NewBuffer(input)
	n.ch, _, err = buf.ReadRune()
	if err != nil {
		return err
	}

	dec = gob.NewDecoder(buf)
	err = dec.Decode(&n.next)
	if err != nil {
		return err
	}

	err = dec.Decode(&n.data)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) GobEncode() ([]byte, error) {
	var buf *bytes.Buffer
	var err error
	var enc *gob.Encoder

	buf = bytes.NewBuffer(make([]byte,32)])
	buf.WriteRune(n.ch)
	enc = gob.NewEncoder(buf)
	err = enc.Encode(n.next)
	if err != nil {
		return buf.Bytes(), err
	}

	err = enc.Encode(n.data)
	if err != nil {
		return buf.Bytes(), err
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
               n2.next = []Node{Node{ch: c}}
            } else {
               if i==len(n2.next) {
                  n2.next = append(n2.next, Node{ch: c})
               } else {
                  n2.next = append(n2.next, Node{})
                  copy(n2.next[i+1:],n2.next[i:])
                  n2.next[i].ch = c
               }
            }
            n2 = &n2.next[i]
         } else {
            n2.next = []Node{Node{ch: c}}
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

