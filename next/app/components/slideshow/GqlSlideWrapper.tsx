import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { GqlTutorialTitle } from "../tutorialtitle/GqlTutorialTitleSlide";
import styles from "./GqlSlideWrapper.module.css";

const fragmentDefinition = graphql(`
  fragment GqlSlideWrapper on SlideWrapper {
    slide {
      # if you forget this, the resulting fragment will have __typename = undefined
      __typename
      #
      # for each slide type
      #
      ... on TutorialTitleSlide {
        ...GqlTutorialTitleSlide
      }
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlSlideWrapper(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const slide = fragment.slide;

  if (!slide.__typename) {
    throw new Error(
      "__typename got undefined - define __typename in GraphQL fragment/query"
    );
  }

  switch (slide.__typename) {
    case "TutorialTitleSlide":
      return (
        <div className={styles.component}>
          <GqlTutorialTitle fragment={slide} />
        </div>
      );
    default:
      return <>wrooong!!!!! implement a new slide type!!</>;
  }
}
