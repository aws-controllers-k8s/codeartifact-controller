# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Helper functions for Code Artifact Domain e2e tests
"""

import logging

class CodeArtifactValidator:
    def __init__(self, codeartifact_client):
        self.codeartifact_client = codeartifact_client

    def get_domain(self, domain_name: str):
        try:
            response = self.codeartifact_client.describe_domain(
                domain=domain_name
            )
            return response["domain"]
        except Exception as e:
            logging.error(f"Error: {e}")
            return None
    
    def domain_exists(self, domain_name: str):
        return self.get_domain(domain_name) is not None
    
    def get_package_group(self, domain_name: str, package_group: str):
        try:
            response = self.codeartifact_client.describe_package_group(
                domain=domain_name,
                packageGroup=package_group,
            )
            return response["packageGroup"]
        except Exception as e:
            logging.error(f"Error: {e}")
            return None
    
    def package_group_exists(self, domain_name: str, package_group: str):
        return self.get_package_group(domain_name, package_group) is not None
    
    def get_resource_tags(self, arn: str):
        try:
            response = self.codeartifact_client.list_tags_for_resource(
                resourceArn=arn
            )
            if len(response["tags"]) == 0:
                return None
            return response["tags"]
        except Exception as e:
            logging.error(f"Error: {e}")
            return None