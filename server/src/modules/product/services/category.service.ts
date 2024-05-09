import { DefaultServiceClass } from '../../../lib/class/default-service.class';
import { CreateCategoryDtoType } from '../dto/create-category.dto';

export class CategoryService extends DefaultServiceClass {
  public async createCategory(dto: CreateCategoryDtoType) {
    return this.prisma.categoriesProduct.create({ data: dto });
  }
}
