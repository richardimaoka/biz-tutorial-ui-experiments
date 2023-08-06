import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { MarkdownView } from "../markdown/MarkdownView";
import { ColumnContentsPosition } from "./ColumnContentsPosition";

const fragmentDefinition = graphql(`
  fragment MarkdownColumnFragment on MarkdownColumn {
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

  if (!fragment.description) return <></>;

  return (
    <ColumnContentsPosition position={position}>
      <MarkdownView fragment={fragment.description} />
    </ColumnContentsPosition>
  );
};
