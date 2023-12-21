import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { MarkdownDefaultStyle } from "../markdown/server-component/MarkdownDefaultStyle";
import { GqlSectionTitleSlide } from "../sectiontitle/GqlSectionTitleSlide";
import { GqlTutorialTitle } from "../tutorialtitle/GqlTutorialTitleSlide";
import styles from "./GqlSlideWrapper.module.css";
import { GqlImageSlide } from "../image/GqlImageSlide";

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
      ... on MarkdownSlide {
        markdownBody
      }
      ... on SectionTitleSlide {
        ...GqlSectionTitleSlide
      }
      ... on ImageSlide {
        ...GqlImageSlide
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
    case "MarkdownSlide":
      return (
        <div className={styles.component}>
          <MarkdownDefaultStyle markdownBody={slide.markdownBody} />
        </div>
      );
    case "SectionTitleSlide":
      return (
        <div className={styles.component}>
          <GqlSectionTitleSlide fragment={slide} />
        </div>
      );
    case "ImageSlide":
      return (
        <div className={styles.component}>
          <GqlImageSlide fragment={slide} />
        </div>
      );
    default:
      return <>wrooong!!!!! implement a new slide type!!</>;
  }
}
