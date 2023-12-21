import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlSlideshowComponent.module.css";
import { GqlSlideWrapper } from "./GqlSlideWrapper";

const fragmentDefinition = graphql(`
  fragment GqlSlideshowComponent on Page {
    slide {
      ...GqlSlideWrapper
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlSlideshowComponent(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.component}>
      {/* contents part */}
      <div className={styles.contents}>
        {fragment.slide && <GqlSlideWrapper fragment={fragment.slide} />}
      </div>
    </div>
  );
}
