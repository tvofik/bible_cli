package themes

	import (
		"fmt"
		"strings"
		"github.com/charmbracelet/lipgloss"
		"golang.org/x/term"
	)


// Styling
var (	
	leftBorder = func() lipgloss.Style{
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0,1)
	}()
	rightBorder = func() lipgloss.Style{
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0,1)
	}()
	border = func() lipgloss.Style{
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		b.Left = "┤"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0,1)
	}()
	nonBorderBodyStyle = border.Copy().UnsetBorderStyle()
	// borderBodyStyle = nonBorderBodyStyle.Copy().BorderStyle(lipgloss.RoundedBorder())
	borderBodyStyle = nonBorderBodyStyle.Copy().Border(lipgloss.RoundedBorder(),false, true)
	

	// Status Bar Theme.
	statusBarStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})


	statusStyle = lipgloss.NewStyle().
		Inherit(statusBarStyle).
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#FF5F87")).
		Padding(0,1)
			
	encodingStyle = statusStyle.Copy().
		Background(lipgloss.Color("#A550DF")).
		Align(lipgloss.Right)
	
	historyStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#874BFD")).
		Padding(0, 1)

	//Default Theme
	defaultHeaderStyle = lipgloss.NewStyle().Padding(0,1)
	defaultBodyStyle = lipgloss.NewStyle().Padding(0,1)
)


// Get the width of terminal window
func getWidth() int {
	width, _, err := term.GetSize(0) // Get the Width and height of the terminal window
	if err != nil{
		return 0 // Error
	}
	return width
}

func Theme1(header, body, abbv string){
	abbv = strings.ToTitle(abbv)
	width := getWidth()
	nonBorderBodyStyle.Width(width)

	title := leftBorder.Render(header)
	footer := rightBorder.Render(abbv)
	
	titleLines := strings.Repeat("─",width - lipgloss.Width(title))
	footerLines := strings.Repeat("─",width - lipgloss.Width(footer))
	
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,title,titleLines))
	fmt.Println(nonBorderBodyStyle.Render(body))
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,footerLines,footer))
}
func Theme2(header, body, abbv string){
	abbv = strings.ToTitle(abbv)
	width := getWidth()


	nonBorderBodyStyle.Width(width-2)

	title := border.Render(header)
	footer := border.Render(abbv)
	
	titleLines := strings.Repeat("─",width - lipgloss.Width(title) - 4)
	footerLines := strings.Repeat("─",width - lipgloss.Width(footer) - 4)
	
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,"╭─",title,titleLines,"─╮"))
	fmt.Println(nonBorderBodyStyle.Render(body))
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,"╰─",footerLines,footer,"─╯"))
}


func Theme3(header, body, abbv string){
	abbv = strings.ToTitle(abbv)
	width := getWidth()
	borderBodyStyle.Width(width-2)

	title := border.Render(header)
	footer := border.Render(abbv)
	
	titleLines := strings.Repeat("─",width - lipgloss.Width(title) - 4)
	footerLines := strings.Repeat("─",width - lipgloss.Width(footer) - 4)
	
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,"╭─",title,titleLines,"─╮"))
	fmt.Println(borderBodyStyle.Render(body))
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,"╰─",footerLines,footer,"─╯"))
}

func Theme4(header, body, abbv string){
	abbv = strings.ToTitle(abbv)
	width := getWidth()
	borderBodyStyle.Width(width-2)

	title := statusBarStyle.Render(header)
	footer := statusBarStyle.Render(abbv)
	
	titleLines := strings.Repeat("─",width - lipgloss.Width(title) - 4)
	footerLines := strings.Repeat("─",width - lipgloss.Width(footer) - 4)
	
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,"╭─",title,titleLines,"─╮"))
	fmt.Println(borderBodyStyle.Render(body))
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,"╰─",footerLines,footer,"─╯"))
}

func Theme5(header, body, abbv string){
	abbv = strings.ToTitle(abbv)
	width := getWidth()
	borderBodyStyle.Width(width-2)

	statusBarStyle.Margin(0,1)

	title := statusBarStyle.Render(header)
	footer := statusBarStyle.Render(abbv)
	
	titleLines := strings.Repeat("─",width - lipgloss.Width(title) - 4)
	footerLines := strings.Repeat("─",width - lipgloss.Width(footer) - 4)
	
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,"╭─",title,titleLines,"─╮"))
	fmt.Println(borderBodyStyle.Render(body))
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,"╰─",footerLines,footer,"─╯"))
}

func Theme6(header, body, abbv string){
	abbv = strings.ToTitle(abbv)
	width := getWidth()
	historyStyle.Width(width)

	title := statusStyle.Render(header)
	footer := encodingStyle.Render(abbv)
	statusVal := statusBarStyle.
			Width(width - lipgloss.Width(title) - lipgloss.Width(footer)).
			Render("")
	
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,title,statusVal,footer))
	fmt.Println(historyStyle.Render(body))
}


func DefaultTheme(header, body, abbv string){
	abbv = strings.ToTitle(abbv)
	width := getWidth()
	defaultBodyStyle.Width(width)
	
	title := defaultHeaderStyle.Render(header)
	footer := defaultHeaderStyle.Render(abbv)	

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center,title,"-",footer))
	fmt.Println(defaultBodyStyle.Render(body))
}
