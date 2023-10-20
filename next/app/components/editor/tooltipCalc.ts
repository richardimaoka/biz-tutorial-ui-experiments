function numLines(tooltipHeight: number, lineHeight: number) {
  // e.g. tooltipHeight = 405, lineHeight = 20, then numLines = 21
  return Math.ceil(tooltipHeight / lineHeight);
}
