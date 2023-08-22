import { FragmentType, graphql, useFragment } from "@/libs/gql";

import styles from "./style.module.css";

const fragmentDefinition = graphql(`
  fragment YouTube_Fragment on YouTubeEmbed {
    embedUrl
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
  if (!fragment.embedUrl || !fragment.width || !fragment.height) {
    return <></>;
  }

  return (
    <iframe
      width={fragment.width}
      height={fragment.height}
      src={fragment.embedUrl}
      title="YouTube video player"
      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
      allowFullScreen
    />
  );
};
