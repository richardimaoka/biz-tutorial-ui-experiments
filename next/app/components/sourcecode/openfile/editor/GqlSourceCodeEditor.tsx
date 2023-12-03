import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { SourceCodeEditor } from "./SourceCodeEditor";

const fragmentDefinition = graphql(`
  fragment GqlSourceCodeEditor on OpenFile {
    content
    oldContent
    language
    editSequence {
      id
      edits {
        text
        range {
          startLineNumber
          startColumn
          endLineNumber
          endColumn
        }
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
}

/**
 * GraphQL-based component calling monaco-editor React component.
 * The purpose of this component is to translate GraphQL fragment
 * into props of the monaco-editor React component.
 */
export function GqlSourceCodeEditor(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  // edit sequence props
  const editSequence = fragment.editSequence
    ? {
        id: fragment.editSequence.id,
        edits: fragment.editSequence.edits ? fragment.editSequence.edits : [],
      }
    : undefined;

  // editor text
  const oldContent = fragment.oldContent ? fragment.oldContent : "";
  const currentContent = fragment.content ? fragment.content : "";

  const doAnimate = editSequence; // editSequence exists
  const editorText = doAnimate ? currentContent : oldContent;

  // editor language
  const language = fragment.language ? fragment.language : "";

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
