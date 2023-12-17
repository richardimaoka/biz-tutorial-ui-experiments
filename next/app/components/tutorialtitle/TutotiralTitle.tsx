import Image from "next/image";
import styles from "./TutorialTitle.module.css";

interface ImageProps {
  src: string;
  width: number;
  height: number;
  caption?: string;
}

interface Props {
  title: string;
  images?: ImageProps[];
}

export function TutorialTitle(props: Props) {
  return (
    <div className={styles.component}>
      {/* text title */}
      <h1 className={styles.title}>{props.title}</h1>

      {/* optional images */}
      {props.images && (
        <div className={styles.imageContainer}>
          {props.images.map((i) => (
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
