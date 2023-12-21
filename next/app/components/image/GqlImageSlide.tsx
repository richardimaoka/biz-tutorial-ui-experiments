import { FragmentType, graphql, useFragment } from "@/libs/gql";
import Image from "next/image";
import styles from "./GqlImageSlide.module.css";

const fragmentDefinition = graphql(`
  fragment GqlImageSlide on ImageSlide {
    image {
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

export function GqlImageSlide(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const image = fragment.image;

  return (
    <div className={styles.component}>
      <Image
        style={{ display: "block" }}
        src={image.src}
        width={image.width}
        height={image.height}
        alt={image.caption ? image.caption : "title image"}
      />
      {image.caption && <div className={styles.caption}>{image.caption}</div>}
    </div>
  );
}
