import { EditorEditable } from "@/app/components/sourcecode2/editor/EditorEditable";
import { FragmentType, graphql, useFragment } from "@/libs/gql";

interface Props {
  editorText: string;
  language: string;
}

export function Interactive(props: Props) {
  return (
    <EditorEditable
      editorText={props.editorText}
      language={props.language}
      edits={[
        {
          range: {
            startLineNumber: 130,
            endLineNumber: 130,
            startColumn: 1,
            endColumn: 20,
          },
          text: null,
        },
      ]}
    />
  );
}
