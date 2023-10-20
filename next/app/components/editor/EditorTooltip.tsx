import { MarkdownDefaultStyle } from "../markdown2/MarkdownDefaultStyle";
import { EditorTooltipCC } from "./EditorTooltipCC";

interface Props {
  markdownBody: string;
  hidden?: boolean;
}

export type EditorTooltipProps = Props;

export function EditorTooltip(props: Props) {
  return (
    <EditorTooltipCC>
      <MarkdownDefaultStyle markdownBody={props.markdownBody} />
    </EditorTooltipCC>
  );
}
