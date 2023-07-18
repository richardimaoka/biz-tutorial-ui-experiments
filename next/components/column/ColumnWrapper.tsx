import { css } from "@emotion/react";
import { dark1MainBg } from "../../libs/colorTheme";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ImageDescriptionColumn } from "./ImageDescriptionColumn";
import { ReactNode } from "react";
import { ImageDescriptionColumnPosition } from "../../libs/gql/graphql";

const fragmentDefinition = graphql(`
  fragment ColumnWrapperFragment on ColumnWrapper {
    column {
      ... on ImageDescriptionColumn {
        ...ImageDescriptionColumnFragment
        position
      }
    }
  }
`);

export interface ColumnWrapperProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

const ColumnPositioning = ({
  position,
  children,
}: {
  position: ImageDescriptionColumnPosition;
  children: ReactNode;
}) => {
  const justifyContent = (p: ImageDescriptionColumnPosition): string => {
    switch (p) {
      case "TOP":
        return "flex-start";
      case "CENTER":
        return "center";
      case "BOTTOM":
        return "flex-end";
    }
  };

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

        display: flex;
        flex-direction: column;
        justify-content: ${justifyContent(position)};
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
      const position = fragment.column?.position
        ? fragment.column?.position
        : "TOP";
      return (
        <ColumnPositioning position={position}>
          <ImageDescriptionColumn fragment={fragment.column} />
        </ColumnPositioning>
      );
    default:
      return <>no matching column</>;
  }
};
