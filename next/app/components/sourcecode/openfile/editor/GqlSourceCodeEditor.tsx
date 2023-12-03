// The component of this file needs to be a SERVER component,
// See the comment below, at the function signature of the component function

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { EditorTooltip } from "../tooltip/EditorTooltip";
import { SourceCodeAnimationHandler } from "./SourceCodeAnimationHandler";

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
  defaultFocusColumn?: string;
}

/**
 * GraphQL-based component calling monaco-editor React component.
 * The purpose of this component is to translate GraphQL fragment
 * into props of the monaco-editor React component.
 *
 * *** CAUTION ***
 * The component of this file needs to be a SERVER component,
 * since the below calls <EditorTooltip>, a SERVER component which calls async/await internally
 * for makrdown processing
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

  const oldContents = fragment.oldContent ? fragment.oldContent : "";
  console.log("oldContents", oldContents);
  const currentContents = fragment.content ? fragment.content : "";

  // editor language
  const language = fragment.language ? fragment.language : "";

  // tooltip props
  const tooltip = fragment.tooltip
    ? {
        lineNumber: fragment.tooltip.lineNumber,
        //default timing is "END", as there could be edits
        timing: fragment.tooltip.timing ? fragment.tooltip.timing : "END",
        children: (
          // Convert the markdownBody string into a React component
          <EditorTooltip markdownBody={fragment.tooltip.markdownBody} />
        ),
      }
    : undefined;

  return (
    <SourceCodeAnimationHandler
      currentContents={currentContents}
      oldContents={oldContents}
      language={language}
      editSequence={editSequence}
      tooltip={tooltip}
      defaultFocusColumn={props.defaultFocusColumn}
    />
  );
}
