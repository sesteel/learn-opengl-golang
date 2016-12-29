package sections

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/raedatoui/learn-opengl-golang/utils"
	"strings"
	"errors"
)

type TitleSlide struct {
	BaseSlide
	font *utils.Font
	Name  string
	lines []string
}

func (s *TitleSlide) Init(a ...interface{}) error {
	f, ok := a[0].(*utils.Font)
	if  ok == false {
		return errors.New("first argument isnt a font")
	}
	s.font = f

	c, ok := a[1].(utils.ColorA)
	if  ok == false {
		return errors.New("second argument isnt a ColorA")
	}
	s.Color = c

	n, ok := a[2].(string)
	if  ok == false {
		return errors.New("second argument isnt a ColorA")
	}
	s.Name = n

	if strings.Contains(s.Name, "\n") {
		s.lines = strings.Split(s.Name, "\n")
	} else {
		s.lines = []string{s.Name}
	}

	return nil
}

func (s *TitleSlide) InitGL() error {
	return nil
}

func (s *TitleSlide) Update() {

}

func (s *TitleSlide) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(s.Color.R, s.Color.G, s.Color.B, s.Color.A)

	s.font.SetColor(1.0, 1.0, 1.0, 1.0)
	for i := 0; i < len(s.lines); i++ {
		s.font.Printf(30, 200+60*float32(i), 1, s.lines[i])

	}
}

func (s *TitleSlide) Close() {

}
