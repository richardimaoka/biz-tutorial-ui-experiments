import { FragmentType, graphql, useFragment } from "@/libs/gql";
import Image from "next/image";
import styles from "./GqlBrowser.module.css";

const fragmentDefinition = graphql(`
  fragment GqlBrowser on Browser {
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

export function GqlBrowser(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const image = fragment.image;

  return (
    <Image
      className={styles.image}
      src={image.src}
      width={image.width}
      height={image.height}
      alt={image.caption ? image.caption : "browser image"}
    />
  );
}
