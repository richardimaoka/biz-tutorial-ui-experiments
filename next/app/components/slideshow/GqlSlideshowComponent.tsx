import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlSlideshowComponent.module.css";
import { GqlSlideWrapper } from "./slides/wrapper/GqlSlideWrapper";
import { GqlModalComponent } from "../modal/GqlModalComponent";
import { GqlNavigation } from "../navigation/GqlNavigation";

const fragmentDefinition = graphql(`
  fragment GqlSlideshowComponent on Page {
    modal {
      ...GqlModalComponent
    }
    slide {
      ...GqlSlideWrapper
    }
    ...GqlNavigation
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
      {fragment && <GqlNavigation fragment={fragment} toInitial />}
    </div>
  );
}
