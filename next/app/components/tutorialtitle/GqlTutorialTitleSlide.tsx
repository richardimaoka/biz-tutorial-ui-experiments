import { FragmentType, graphql, useFragment } from "@/libs/gql";
import Image from "next/image";
import styles from "./GqlTutorialTitleSlide.module.css";

const fragmentDefinition = graphql(`
  fragment GqlTutorialTitle on TutorialTitleSlide {
    title
    images {
      src
      width
      height
      caption
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlTutorialTitle(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  return (
    <div className={styles.component}>
      {/* text title */}
      <h1 className={styles.title}>{fragment.title}</h1>

      {/* optional images */}
      {fragment.images && (
        <div className={styles.imageContainer}>
          {fragment.images.map((i) => (
            <div className={styles.image} key={i.src}>
              <Image
                style={{ display: "block" }}
                src={i.src}
                width={i.width}
                height={i.height}
                alt={i.caption ? i.caption : "title image"}
              />
              {i.caption && <div className={styles.caption}>{i.caption}</div>}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
