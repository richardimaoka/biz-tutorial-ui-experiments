import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ButtonToInitialStep } from "./ButtonToInitialStep";
import { GqlColumnTabs } from "./GqlColumnTabs";
import styles from "./GqlTutorialHeader.module.css";

const fragmentDefinition = graphql(`
  fragment GqlTutorialHeader on Page2 {
    ...GqlColumnTabs
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  selectTab: string;
}

export function GqlTutorialHeader(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.component}>
      <GqlColumnTabs fragment={fragment} selectTab={props.selectTab} />
      <ButtonToInitialStep href="" />
    </div>
  );
}
