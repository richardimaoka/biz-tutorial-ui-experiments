import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ColumnHeader } from "./ColumnHeader";

import styles from "./style.module.css";
import { nonNullArray } from "@/libs/nonNullArray";
import { ModalFrame } from "../modal/ModalFrame";
import { Navigation } from "../navigation/Navigation";
import { Carousel } from "./Carousel";

const fragmentDefinition = graphql(`
  fragment VisibleColumn_Fragment on Page {
    ...ColumnHeader_Fragment
    ...Carousel_Fragment
    columns {
      ...ColumnWrapperComponent_Fragment
      name
    }
    modal {
      ...ModalFrameFragment
    }
    ...Navigation_Fragment
  }
`);

interface VisibleColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  step: string;
  selectColumn?: string;
  openFilePath?: string;
  skipAnimation?: boolean;
  autoNextSeconds?: number;
}

export const VisibleColumn = (props: VisibleColumnProps) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment?.columns) {
    return <div></div>;
  }
  const columns = nonNullArray(fragment.columns);

  const selectColumn = props.selectColumn
    ? decodeURI(props.selectColumn)
    : columns.length > 0 && columns[0].name
    ? columns[0].name
    : "";

  // const visibleColumn = columns.find((column) => column.name === selectColumn);

  return (
    <div className={styles.visiblecolumn}>
      <ColumnHeader
        fragment={fragment}
        selectColumn={selectColumn}
        openFilePath={props.openFilePath}
        step={props.step}
      />
      <div className={styles.body}>
        {/* above <div> + .body style is necessary to control the height of visible column = 100svh */}
        {fragment.modal ? (
          <ModalFrame fragment={fragment.modal}>
            <Carousel
              fragment={fragment}
              step={props.step}
              skipAnimation={props.skipAnimation}
            />
          </ModalFrame>
        ) : (
          <Carousel
            fragment={fragment}
            step={props.step}
            skipAnimation={props.skipAnimation}
          />
        )}
      </div>
      <div className={styles.button}>
        <Navigation
          fragment={fragment}
          autoNextSeconds={props.autoNextSeconds}
        />
      </div>
    </div>
  );
};
