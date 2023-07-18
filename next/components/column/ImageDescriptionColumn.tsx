import Image from "next/image";
import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ImageDescriptionOrder } from "../../libs/gql/graphql";
import { MarkdownView } from "../markdown/MarkdownView";
import { ImageCentered } from "../image/ImageCentered";

const fragmentDefinition = graphql(`
  fragment ImageDescriptionColumnFragment on ImageDescriptionColumn {
    description {
      ...MarkdownFragment
    }
    image {
      ...ImageCenteredFragment
    }
    order
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

  switch (order) {
    case "DESCRIPTION_THEN_IMAGE":
      return (
        <>
          {fragment.description && (
            <MarkdownView fragment={fragment.description} />
          )}
          {fragment.image && <ImageCentered fragment={fragment.image} />}
        </>
      );
    case "IMAGE_THEN_DESCRIPTION":
      return (
        <>
          {fragment.image && <ImageCentered fragment={fragment.image} />}
          {fragment.description && (
            <MarkdownView fragment={fragment.description} />
          )}
        </>
      );
  }
};
