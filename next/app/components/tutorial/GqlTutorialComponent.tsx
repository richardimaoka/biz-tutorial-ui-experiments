import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { Carousel } from "../carousel/Carousel";
import styles from "./GqlTutorialComponent.module.css";
import { columnWidthPx } from "./definitions";
import { GqlTutorialHeader } from "./header/GqlTutorialHeader";
import { GqlColumnWrappers } from "./column/GqlColumnWrappers";

const fragmentDefinition = graphql(`
  fragment GqlTutorialComponent on Page2 {
    ...GqlTutorialHeader
    ...GqlColumnWrappers
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  selectTab: string;
}

export function GqlTutorialComponent(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.component}>
      {/* header part */}
      <div className={styles.header}>
        <GqlTutorialHeader fragment={fragment} selectTab={props.selectTab} />
      </div>
      {/* contents part */}
      <div className={styles.contents}>
        <Carousel currentIndex={0} columnWidth={columnWidthPx}>
          <GqlColumnWrappers fragment={fragment} />
        </Carousel>
      </div>
    </div>
  );
}
