package main

type Part struct {
	X, Y int
}

type SnakeBody struct {
	Parts          []Part
	Xspeed, Yspeed int
}

func (sb *SnakeBody) ChangeDir(vertical, horizontal int) {
	sb.Xspeed = horizontal
	sb.Yspeed = vertical
}

func (sb *SnakeBody) Update(width, height int, longerSnake bool) {
	sb.Parts = append(
		sb.Parts,
		sb.Parts[len(sb.Parts)-1].GetUpdatedPart(
			sb,
			width,
			height,
		),
	)
	if !longerSnake {
		sb.Parts = sb.Parts[1:]
	}
}

func (sb *SnakeBody) ResetPos(width, height int) {
	snakeParts := []Part{
		{
			X: int(width / 2),
			Y: int(height / 2),
		},
		{
			X: int(width/2) + 1,
			Y: int(height / 2),
		},
		{
			X: int(width/2) + 2,
			Y: int(height / 2),
		},
	}
	sb.Parts = snakeParts
	sb.Xspeed = 1
	sb.Yspeed = 0
}

func (sp *Part) GetUpdatedPart(sb *SnakeBody, width, height int) Part {
	newPart := *sp
	newPart.X = (newPart.X + sb.Xspeed) % width
	if newPart.X < 0 {
		newPart.X += width
	}
	newPart.Y = (newPart.Y + sb.Yspeed) % height
	if newPart.Y < 0 {
		newPart.Y += height
	}
	return newPart
}
