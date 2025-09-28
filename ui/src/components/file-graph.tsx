"use client";

import useEdges from "@/hooks/use-edges";
import useFiles from "@/hooks/use-files";
import { usePathname } from "next/navigation";
import { useEffect, useRef } from "react";
import * as vis from "vis-network";

const FileGraph = () => {
  const graphRef = useRef(null);
  const pathname = usePathname();
  const { files } = useFiles();
  const edgesData = useEdges();
  const course = pathname.split("/")[1];

  const nodes: vis.Node[] = files
    ?.map((f) => {
      const parts = f.file_path.split("/");
      const rootIdx = parts.indexOf("school");
      const label = parts.slice(rootIdx + 1).join("/");
      return {
        id: f.file_path,
        label: label,
      };
    })
    .filter((f) => !f.label.includes("DS_Store"));

  const edges = edgesData?.map((e) => e as vis.Edge);

  useEffect(() => {
    if (!graphRef.current) return;
    const network = new vis.Network(
      graphRef.current,
      { nodes, edges },
      options,
    );
    network.on("beforeDrawing", (ctx: CanvasRenderingContext2D) => {
      if (course === "") return;
      nodes.forEach((node) => {
        const box = network.getBoundingBox(node.id!);
        if (node.label!.includes(course)) {
          ctx.fillStyle = "#fff085";
          ctx.roundRect(
            box.left,
            box.top,
            box.right - box.left,
            box.bottom - box.top,
            10,
          );
          ctx.fill();
        }
      });
    });
    network.redraw();
  }, [course, nodes, edges]);

  return <div ref={graphRef} className="border-l basis-1/3"></div>;
};

export default FileGraph;

const options: vis.Options = {
  nodes: {
    value: 10,
    borderWidth: 1,
    color: {
      border: "#dadada",
      background: "#ffffff",
    },
    shape: "box",
    font: {
      face: "Geist Mono",
    },
    scaling: {
      min: 10,
      max: 30,
    },
  },
  edges: {
    color: {
      color: "#3b82f6",
    },
  },
  interaction: {
    hover: true,
  },
};
