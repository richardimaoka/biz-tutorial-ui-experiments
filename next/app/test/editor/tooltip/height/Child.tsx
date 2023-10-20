import { EditorTooltip } from "@/app/components/editor/tooltip/EditorTooltip";

interface Props {
  markdownBody: string;
}

export function Child(props: Props) {
  return <div>{<EditorTooltip markdownBody={props.markdownBody} />}</div>;
}
