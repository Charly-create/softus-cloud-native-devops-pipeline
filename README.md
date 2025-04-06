# 🚀 DevOps Pipeline by Charly-create

**GitHub:** [github.com/Charly-create](https://github.com/Charly-create)  
**Tech Stack:**  
- AWS EKS (Terraform)  
- GitHub Actions (CI/CD)  
- Trivy (Security)  
- Prometheus (Monitoring)  

## How It Works
1. **Push to `main` branch** → Triggers GitHub Actions  
2. **Build & Scan** → Containerize app + Trivy scan  
3. **Provision EKS** → Terraform creates the cluster  
4. **Deploy** → App + monitoring stack are deployed  

## Setup
```bash
# Clone repo
git clone https://github.com/Charly-create/cloud-native-devops-pipeline.git

# Deploy (via GitHub Actions)
git push origin main