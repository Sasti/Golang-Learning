package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addNumbersTestify(t *testing.T) {
	sum, err := addNumbers("3", "3")

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, "6", sum, "they should be equal")
}

func Test_addNumbers(t *testing.T) {
	sum, err := addNumbers("3", "3")

	if err != nil {
		t.Error(err)
	}

	if sum != "6" {
		t.Log("Didn't get expected 6 from function")
		t.Fail()
	}
}

func Test_addNumbersWithDataProvider(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Simple add",
			args{
				nums: []string{"2", "3"},
			},
			"5",
			false,
		},
		{
			"Big numbers",
			args{
				nums: []string{"20000000000", "30000000000"},
			},
			"50000000000",
			false,
		},
		{
			"Add decimal numbers",
			args{
				nums: []string{"20.5", "30.2"},
			},
			"50.7",
			false,
		},
		{
			"Not enough arguments",
			args{
				nums: []string{"2"},
			},
			"",
			true,
		},
		{
			"Argument is not a number",
			args{
				nums: []string{"2", "fasdf5"},
			},
			"",
			true,
		},
	}

	// Give the hint that this test can be executed in parallel
	t.Parallel()

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addNumbers(tt.args.nums...)
			if (err != nil) != tt.wantErr {
				t.Errorf("addNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("addNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
