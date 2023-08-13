import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { MarkdownView } from "../markdown/MarkdownView";

const fragmentDefinition = graphql(`
  fragment MarkdownColumn_Fragment on MarkdownColumn {
    description {
      ...MarkdownFragment
    }
    contentsPosition
  }
`);

export interface MarkdownColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const MarkdownColumn = (props: MarkdownColumnProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const position = fragment.contentsPosition
    ? fragment.contentsPosition
    : "TOP";

  console.log(fragment);

  if (!fragment.description) return <></>;

  return <MarkdownView fragment={fragment.description} />;
};
