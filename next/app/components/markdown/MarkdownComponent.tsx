interface Props {
  markdownBody: string;
}

export function MarkdownComponent(props: Props) {
  return <div data-testid="markdown">{props.markdownBody}</div>;
}
