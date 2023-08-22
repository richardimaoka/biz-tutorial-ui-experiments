import { FragmentType, graphql, useFragment } from "@/libs/gql";

import styles from "./style.module.css";
import { YouTubeView } from "./YouTubeView";

const fragmentDefinition = graphql(`
  fragment YouTubeColumn_Fragment on YouTubeColumn {
    youtube {
      ...YouTube_Fragment
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const YouTubeColumn = (props: Props) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.column}>
      {fragment.youtube && <YouTubeView fragment={fragment.youtube} />}
    </div>
  );
};
