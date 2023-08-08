import { ColumnHeader } from "./ColumnHeader";

import styles from "./style.module.css";
import { nonNullArray } from "@/libs/nonNullArray";
import { ModalComponent } from "./modal/ModalComponent";
import { Navigation } from "../navigation/Navigation";
import { Carousel } from "./Carousel";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment VisibleColumn_Fragment on Page {
    ...ColumnHeader_Fragment
    ...Carousel_Fragment
    columns {
      ...ColumnWrapperComponent_Fragment
      name
    }
    modal {
      ...ModalComponentFragment
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
        {fragment.modal && <ModalComponent fragment={fragment.modal} />}
        <Carousel
          fragment={fragment}
          step={props.step}
          skipAnimation={props.skipAnimation}
        />
      </div>
      <div className={styles.bottom}>
        <Navigation
          fragment={fragment}
          autoNextSeconds={props.autoNextSeconds}
        />
      </div>
    </div>
  );
};
