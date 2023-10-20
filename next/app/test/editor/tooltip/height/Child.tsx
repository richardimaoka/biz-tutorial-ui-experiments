import { EditorTooltip } from "@/app/components/editor/EditorTooltip";

interface Props {
  markdownBody: string;
}

export function Child(props: Props) {
  return <div>{<EditorTooltip markdownBody={props.markdownBody} />}</div>;
}
