package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewParking(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *Parking
	}{
		{
			name: "create parking lot success",
			args: args{
				4,
			},
			want: &Parking{
				Available:        &IntMinHeap{1, 2, 3, 4},
				Occupied:         map[int]Car{},
				slotByPlatNumber: map[string]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parking := NewParking(tt.args.n)
			assert.Equal(t, tt.want, parking)
		})
	}
}

func TestParking_Park(t *testing.T) {
	type fields struct {
		available *IntMinHeap
		occupied  map[int]Car
		slotByReg map[string]int
	}
	type args struct {
		plateNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "park available",
			fields: fields{
				available: &IntMinHeap{
					2, 3,
				},
				occupied: map[int]Car{
					1: {
						"test-1",
						"red",
					},
				},
				slotByReg: map[string]int{
					"test-1": 1,
				},
			},
			args: args{
				plateNumber: "test-2",
			},
			want: 2,
		},
		{
			name: "park available nearest",
			fields: fields{
				available: &IntMinHeap{
					1, 3,
				},
				occupied: map[int]Car{
					2: {
						"test-2",
						"red",
					},
				},
				slotByReg: map[string]int{
					"test-2": 2,
				},
			},
			args: args{
				plateNumber: "test-4",
			},
			want: 1,
		},
		{
			name: "park full",
			fields: fields{
				available: &IntMinHeap{},
				occupied: map[int]Car{
					1: {
						"test-1",
						"red",
					},
					2: {
						"test-2",
						"red",
					},
				},
				slotByReg: map[string]int{
					"test-1": 1,
					"test-2": 2,
				},
			},
			args: args{
				plateNumber: "test-3",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parking{
				Available:        tt.fields.available,
				Occupied:         tt.fields.occupied,
				slotByPlatNumber: tt.fields.slotByReg,
			}
			actualSlot, err := p.Park(tt.args.plateNumber)
			if tt.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, fmt.Errorf("sorry, parking lot full").Error())
			}

			assert.Equal(t, tt.want, actualSlot)
		})
	}
}

func TestParking_Leave(t *testing.T) {
	type fields struct {
		available        *IntMinHeap
		occupied         map[int]Car
		slotByPlatNumber map[string]int
	}
	type args struct {
		plateNumber string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		expectedAvailable int
		expectedFreeSlot  int
		wantErr           bool
		expectedErr       error
	}{
		{
			name: "leave park success",
			fields: fields{
				available: &IntMinHeap{
					2, 3,
				},
				occupied: map[int]Car{
					1: {
						"test-1",
						"red",
					},
				},
				slotByPlatNumber: map[string]int{
					"test-1": 1,
				},
			},
			args: args{
				plateNumber: "test-1",
			},
			expectedFreeSlot:  1,
			expectedAvailable: 3,
		},
		{
			name: "car not found in any park slot",
			fields: fields{
				available: &IntMinHeap{
					2, 3,
				},
				occupied: map[int]Car{
					1: {
						"test-1",
						"red",
					},
				},
				slotByPlatNumber: map[string]int{
					"test-1": 1,
				},
			},
			args: args{
				plateNumber: "test-2",
			},
			wantErr:     true,
			expectedErr: fmt.Errorf("car not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parking{
				Available:        tt.fields.available,
				Occupied:         tt.fields.occupied,
				slotByPlatNumber: tt.fields.slotByPlatNumber,
			}
			slot, err := p.Leave(tt.args.plateNumber)
			if tt.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Equal(t, tt.expectedFreeSlot, slot)
				assert.Equal(t, tt.expectedAvailable, p.Available.Len())
			}

		})
	}
}
