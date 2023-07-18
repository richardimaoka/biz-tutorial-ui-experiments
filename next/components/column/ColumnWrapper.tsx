import { css } from "@emotion/react";
import { dark1MainBg } from "../../libs/colorTheme";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ImageDescriptionColumn } from "./ImageDescriptionColumn";
import { ReactNode } from "react";

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

const ColumnPositioning = ({ children }: { children: ReactNode }) => {
  return (
    <div
      css={css`
        @media (max-width: 768px) {
          width: 100vw;
          height: 100vh;
        }
        width: 768px;
        min-height: 100vh;
        overflow: auto;
        background-color: ${dark1MainBg};
      `}
    >
      {children}
    </div>
  );
};

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
      return (
        <ColumnPositioning>
          <ImageDescriptionColumn fragment={fragment.column} />
        </ColumnPositioning>
      );
    default:
      return <>no matching column</>;
  }
};
