import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ImageDescriptionColumn } from "./ImageDescriptionColumn";
import { MarkdownColumn } from "./MarkdownColumn";
import { BackgroundImageColumn } from "./BackgroundImageColumn";
import { TerminalColumn } from "./TerminalColumn";
import { SourceCodeColumn } from "./SourceCodelColumn";
import { ColumnContentsPosition } from "./ColumnContentsPosition";

const fragmentDefinition = graphql(`
  fragment ColumnWrapperFragment on ColumnWrapper {
    column {
      ... on ImageDescriptionColumn {
        ...ImageDescriptionColumnFragment
        contentsPosition
      }
      ... on BackgroundImageColumn {
        ...BackgroundImageColumnFragment
      }
      ... on MarkdownColumn {
        ...MarkdownColumnFragment
      }
      ... on TerminalColumn {
        ...TerminalColumnFragment
      }
      ... on SourceCodeColumn {
        ...SourceCodeColumnFragment
      }
    }
  }
`);

export interface ColumnWrapperProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  step: string;
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
      const position = fragment.column.contentsPosition
        ? fragment.column.contentsPosition
        : "TOP";
      return (
        <ColumnContentsPosition position={position}>
          <ImageDescriptionColumn fragment={fragment.column} />
        </ColumnContentsPosition>
      );
    case "BackgroundImageColumn":
      return <BackgroundImageColumn fragment={fragment.column} />;
    case "MarkdownColumn":
      return <MarkdownColumn fragment={fragment.column} />;
    case "TerminalColumn":
      return (
        <ColumnContentsPosition position="TOP">
          <TerminalColumn fragment={fragment.column} />
        </ColumnContentsPosition>
      );
    case "SourceCodeColumn":
      return (
        <ColumnContentsPosition position="TOP">
          <SourceCodeColumn fragment={fragment.column} step={props.step} />
        </ColumnContentsPosition>
      );
    default:
      return <>no matching column</>;
  }
};
