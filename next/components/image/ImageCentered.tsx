import Image from "next/image";
import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";

const ImageCenteredFragment = graphql(`
  fragment ImageCenteredFragment on ImageCentered {
    width
    height
    path
  }
`);

export interface ImageCenteredViewProps {
  fragment: FragmentType<typeof ImageCenteredFragment>;
}

export const ImageCentered = (props: ImageCenteredViewProps) => {
  const fragment = useFragment(ImageCenteredFragment, props.fragment);

  return fragment && fragment.width && fragment.height && fragment.path ? (
    <Image
      css={css`
        display: block;
        margin: 8px auto;
      `}
      src={fragment.path}
      width={fragment.width}
      height={fragment.height}
      alt={"centered image"}
    />
  ) : (
    <></>
  );
};
