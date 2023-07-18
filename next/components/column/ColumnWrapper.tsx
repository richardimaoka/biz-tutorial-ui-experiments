import Image from "next/image";
import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ImageDescriptionColumn } from "./ImageDescriptionColumn";

const fragmentDefinition = graphql(`
  fragment ColumnWrapperFragment on ColumnWrapper {
    column {
      ... on ImageDescriptionColumn {
        ...ImageDescriptionColumnFragment
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
    default:
      return <>no matching column</>;
  }
};
