import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { SourceCodeEditor } from "./SourceCodeEditor";

const fragmentDefinition = graphql(`
  fragment GqlSourceCodeEditor on OpenFile {
    content
    language
    edits {
      text
      range {
        startLineNumber
        startColumn
        endLineNumber
        endColumn
      }
    }
    tooltip {
      markdownBody
      lineNumber
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  skipAnimation?: boolean;
}

export function GqlSourceCodeEditor(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  // basic editor props
  const editorText = fragment.content ? fragment.content : "";
  const language = fragment.language ? fragment.language : "";

  // edit sequence props
  const edits = fragment.edits ? fragment.edits : [];
  const editSequence =
    edits.length > 0
      ? {
          edits: edits,
          skipAnimation: props.skipAnimation,
        }
      : undefined;

  // tooltip props
  const tooltip = fragment.tooltip
    ? {
        markdownBody: fragment.tooltip.markdownBody,
        lineNumber: fragment.tooltip.lineNumber,
      }
    : undefined;

  return (
    <SourceCodeEditor
      editorText={editorText}
      language={language}
      editSequence={editSequence}
      tooltip={tooltip}
    />
  );
}
