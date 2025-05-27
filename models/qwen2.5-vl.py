from modelscope import Qwen2_5_VLForConditionalGeneration, AutoProcessor
from qwen_vl_utils import process_vision_info
import torch
import argparse
import os
os.environ["PYTORCH_CUDA_ALLOC_CONF"] = "expandable_segments:True"

def analyze_video(video_path: str, prompt: str, fps: float = 1.0, max_pixels: int = 854 * 480):
    """
    优化后的视频分析函数
    参数:
        video_path: 视频文件路径(如"/path/to/video.mp4")
        prompt: 分析指令(如"检测火焰并输出时间戳")
        fps: 视频帧率(默认1.0)
        max_pixels: 最大像素数(默认854 * 480)
    """
    # 初始化模型和处理器
    model = Qwen2_5_VLForConditionalGeneration.from_pretrained(
        "models/Qwen2.5-VL-7B",
        torch_dtype="auto",
        device_map="auto"
    )
    processor = AutoProcessor.from_pretrained(
        "models/Qwen2.5-VL-7B",
        use_fast=True,
        videos_kwargs={  
            "fps": fps,
            "max_pixels": max_pixels,
            "min_pixels": 256 * 28 * 28  
          }
          )

    # 构建输入消息(网页1标准格式)
    messages = [{
        "role": "user",
        "content": [
            {
                "type": "video",
                "video": f"file://{video_path}",
                "max_pixels": max_pixels,
                "fps": fps,
            },
            {"type": "text", "text": prompt}
        ]
    }]

    # 预处理输入(网页6标准流程)
    text = processor.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
    image_inputs, video_inputs = process_vision_info(messages)
    
    # 生成模型输入
    inputs = processor(
        text=[text],
        images=image_inputs,
        videos=video_inputs,
        padding=True,
        return_tensors="pt"
    ).to(model.device)

    # 模型推理
    with torch.no_grad():
        generated_ids = model.generate(
            **inputs,
            max_new_tokens=256
        )
    
    # 结果解码
    outputs = generated_ids[:, inputs.input_ids.shape[1]:]
    return processor.batch_decode(outputs, skip_special_tokens=True)[0]

def main():
    parser = argparse.ArgumentParser(description='视频分析工具')
    parser.add_argument("--video", required=True, help="视频文件路径")
    parser.add_argument("--prompt", required=True, help="分析指令")
    parser.add_argument("--fps", type=float, default=3.0, help="采样帧率(默认3.0)")
    parser.add_argument("--max_pixels", type=int, default=854 * 480, help="最大分辨率(默认854x480)")
    
    args = parser.parse_args()
    
    result = analyze_video(
        video_path=args.video,
        prompt=args.prompt,
        fps=args.fps,
        max_pixels=args.max_pixels
    )
    print(f"视频分析结果:\n {result}")

if __name__ == "__main__":
    main()  
