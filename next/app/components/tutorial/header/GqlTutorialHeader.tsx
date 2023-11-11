import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ButtonToInitialStep } from "./buttons/ButtonToInitialStep";
import { GqlColumnTabs } from "./tabs/GqlColumnTabs";
import styles from "./GqlTutorialHeader.module.css";

const fragmentDefinition = graphql(`
  fragment GqlTutorialHeader on Page {
    ...GqlColumnTabs
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlTutorialHeader(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.component}>
      <GqlColumnTabs fragment={fragment} />
      <ButtonToInitialStep href="" />
    </div>
  );
}
