import { MarkdownDefaultStyle } from "../../markdown2/server-component/MarkdownDefaultStyle";
import { EditorTooltipCC } from "./EditorTooltipCC";

interface Props {
  markdownBody: string;
  hidden?: boolean;
}

export type EditorTooltipProps = Props;

export function EditorTooltip(props: Props) {
  return (
    // Since Markdown component is a server component with async rehype-react,
    // client component needs to interleave with the server component using children-passing
    <EditorTooltipCC>
      <MarkdownDefaultStyle markdownBody={props.markdownBody} />
    </EditorTooltipCC>
  );
}
