import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ImageDescriptionColumn } from "./ImageDescriptionColumn";
import { MarkdownColumn } from "./MarkdownColumn";
import { BackgroundImageColumn } from "./BackgroundImageColumn";

const fragmentDefinition = graphql(`
  fragment ColumnWrapperFragment on ColumnWrapper {
    column {
      ... on ImageDescriptionColumn {
        ...ImageDescriptionColumnFragment
      }
      ... on BackgroundImageColumn {
        ...BackgroundImageColumnFragment
      }
      ... on MarkdownColumn {
        ...MarkdownColumnFragment
      }
    }
  }
`);

export interface ColumnWrapperProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const ColumnWrapper = (props: ColumnWrapperProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const typename = fragment.column?.__typename;

  if (!fragment.column) {
    return <></>;
  }

  if (!typename) {
    return <></>;
  }

  switch (typename) {
    case "ImageDescriptionColumn":
      return <ImageDescriptionColumn fragment={fragment.column} />;
    case "BackgroundImageColumn":
      return <BackgroundImageColumn fragment={fragment.column} />;
    case "MarkdownColumn":
      return <MarkdownColumn fragment={fragment.column} />;
    default:
      return <>no matching column</>;
  }
};
