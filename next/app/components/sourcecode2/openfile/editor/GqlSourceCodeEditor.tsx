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
      timing
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  skipAnimation?: boolean;
}

/**
 * GraphQL-based component calling monaco-editor React component.
 * The purpose of this component is to translate GraphQL fragment
 * into props of the monaco-editor React component.
 */
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
        ...fragment.tooltip,
        timing: fragment.tooltip.timing ? fragment.tooltip.timing : "END", //default timing is end, as there could be edits
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
