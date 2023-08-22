import { FragmentType, graphql, useFragment } from "@/libs/gql";

import styles from "./style.module.css";

const fragmentDefinition = graphql(`
  fragment YouTube_Fragment on YouTubeEmbed {
    videoId
    width
    height
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const YouTubeView = (props: Props) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  // TODO: use error.tsx
  if (!fragment.videoId || !fragment.width || !fragment.height) {
    return <></>;
  }

  return (
    <iframe
      width={fragment.width}
      height={fragment.height}
      src="https://www.youtube.com/embed/xz6aeeeJR-g"
      title="YouTube video player"
      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
      allowFullScreen
    />
  );
};
