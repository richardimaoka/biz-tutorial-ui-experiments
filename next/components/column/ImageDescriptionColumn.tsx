import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ImageDescriptionOrder } from "../../libs/gql/graphql";
import { ImageCentered } from "../image/ImageCentered";
import { MarkdownView } from "../markdown/MarkdownView";
import { ColumnContentsPosition } from "./ColumnContentsPosition";

const fragmentDefinition = graphql(`
  fragment ImageDescriptionColumnFragment on ImageDescriptionColumn {
    description {
      ...MarkdownFragment
    }
    image {
      ...ImageCenteredFragment
    }
    order
    contentsPosition
  }
`);

export interface ImageDescriptionColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const ImageDescriptionColumn = (
  props: ImageDescriptionColumnProps
): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const order: ImageDescriptionOrder = fragment.order
    ? fragment.order
    : "DESCRIPTION_THEN_IMAGE"; // default order
  const position = fragment.contentsPosition
    ? fragment.contentsPosition
    : "TOP";

  console.log("ImageDescriptionColumn", fragment);

  switch (order) {
    case "DESCRIPTION_THEN_IMAGE":
      return (
        <ColumnContentsPosition position={position}>
          {fragment.description && (
            <MarkdownView fragment={fragment.description} />
          )}
          {fragment.image && <ImageCentered fragment={fragment.image} />}
        </ColumnContentsPosition>
      );
    case "IMAGE_THEN_DESCRIPTION":
      return (
        <ColumnContentsPosition position={position}>
          {fragment.image && <ImageCentered fragment={fragment.image} />}
          {fragment.description && (
            <MarkdownView fragment={fragment.description} />
          )}
        </ColumnContentsPosition>
      );
  }
};
