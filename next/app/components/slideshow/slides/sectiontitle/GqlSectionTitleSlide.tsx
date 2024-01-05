import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlSectionTitleSlide.module.css";

const fragmentDefinition = graphql(`
  fragment GqlSectionTitleSlide on SectionTitleSlide {
    title
    sectionNum
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlSectionTitleSlide(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  return (
    <div className={styles.component}>
      {/* text title */}
      <h1 className={styles.title}>
        セクション {fragment.sectionNum}
        <br />
        {fragment.title}
      </h1>
    </div>
  );
}
