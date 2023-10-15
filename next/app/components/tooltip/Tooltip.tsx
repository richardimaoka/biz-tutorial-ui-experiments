import { MarkdownDefaultStyle } from "../markdown/MarkdownDefaultStyle";

interface Props {
  body: string; // can be markdown
}

export function Tooltip(props: Props) {
  return (
    <div>
      <MarkdownDefaultStyle markdownBody={props.body} />
    </div>
  );
}
