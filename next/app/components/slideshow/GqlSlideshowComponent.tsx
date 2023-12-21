import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlSlideshowComponent.module.css";
import { GqlSlideWrapper } from "./GqlSlideWrapper";
import { GqlModalComponent } from "../modal/GqlModalComponent";

const fragmentDefinition = graphql(`
  fragment GqlSlideshowComponent on Page {
    modal {
      ...GqlModalComponent
    }
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
      {fragment.slide && <GqlSlideWrapper fragment={fragment.slide} />}
      {fragment.modal && <GqlModalComponent fragment={fragment.modal} />}
    </div>
  );
}