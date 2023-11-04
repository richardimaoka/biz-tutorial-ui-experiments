import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { SourceCodeEditor } from "./SourceCodeEditor";

const fragmentDefinition = graphql(`
  fragment GqlSourceCodeEditor on OpenFile {
    content
    language
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlSourceCodeEditor(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const editorText = fragment.content ? fragment.content : "";
  const language = fragment.language ? fragment.language : "";

  return <SourceCodeEditor editorText={editorText} language={language} />;
}
